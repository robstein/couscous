#include "Arduino.h"
#include "FastLED.h"
#include <ESP8266WiFi.h>
#include <string.h>

#define HTTP_STATUS_OK                                                         \
  String("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: "          \
         "close\r\n\r\n")
#define HTTP_STATUS_INTERNAL_ERROR                                             \
  String("HTTP/1.1 500 OK\r\nContent-Type: text/html\r\nConnection: "          \
         "close\r\n\r\n")
#define HTTP_STATUS_NOT_FOUND                                                  \
  String("HTTP/1.0 404 Not Found\r\nContent-Type: text/html\r\nConnection: "   \
         "close\r\n\r\n")

const int NUM_LEDS = 760;
CRGBArray<NUM_LEDS> leds;

enum Pattern {
  OFF = 0,
  CYLON = 1,
  SWIRL = 2,
};
// Pattern state = OFF;
String state;

typedef struct {
  int status;
  String body;
} response;

typedef response *(*handleFn)(String, String);

typedef struct {
  String method;
  String path;
  handleFn handle;
} handler;

response *handle_get_config(String headers, String body) {
  Serial.println("[Debug] Received GET config");
  response *r = (response *)malloc(sizeof(response));
  r->status = 200;
  r->body = String(state);
  return r;
}

response *handle_post_config(String headers, String body) {
  Serial.println("[Debug] Received POST config");
  body.trim();
  // state = (Pattern)body.toInt();
  state = body;

  response *r = (response *)malloc(sizeof(response));
  r->status = 200;
  r->body = String(state);

  return r;
}

WiFiServer server(80);
const int HANDLERS_SIZE = 3;
handler handlers[HANDLERS_SIZE] = {
    (handler){"GET", "/config", handle_get_config},
    (handler){"POST", "/config", handle_post_config},
};

const int DATA_PIN = 6;
const int BRIGHTNESS = 96;

void setup() {
  Serial.begin(9600);

  WiFi.mode(WIFI_AP);
  WiFi.softAP("nodemcu-testing", "12345678");
  server.begin();

  FastLED.addLeds<NEOPIXEL, DATA_PIN>(leds, NUM_LEDS)
      .setCorrection(TypicalLEDStrip);
  FastLED.setBrightness(BRIGHTNESS);

  state = "";
}

void ledShow() {}

void serve() {
  WiFiClient client = server.available();
  if (client) {
    handleFn fn = NULL;
    String headers = "";
    String body = "";

    bool scanningfirstLine = true;
    bool scanningHeaders = false;
    bool scanningBody = false;

    while (client.connected()) {
      if (client.available()) {
        String line;
        if (!scanningBody) {
          line = client.readStringUntil('\r');
        } else {
          line = client.readString();
        }
        Serial.print(line);

        if (scanningfirstLine) {
          for (int i = 0; i < HANDLERS_SIZE; i++) {
            if (line.startsWith(handlers[i].method + " " + handlers[i].path)) {
              fn = handlers[i].handle;
              break;
            }
          }
          scanningfirstLine = false;
          scanningHeaders = true;
          continue;
        }

        if (scanningHeaders) {
          headers += line;
          if ((line.length() == 1 && line[0] == '\n')) {
            scanningHeaders = false;
            scanningBody = true;
            continue;
          }
        }

        if (scanningBody) {
          body += line;
          break;
        }
      }
    }

    if (fn != NULL) {
      response *r = (fn)(headers, body);

      if (r->status == 200) {
        client.print(HTTP_STATUS_OK + r->body);
      } else {
        client.print(HTTP_STATUS_INTERNAL_ERROR);
      }
    } else {
      client.print(HTTP_STATUS_NOT_FOUND);
    }

    delay(1);
    client.stop();
  }
}

void loop() {
  ledShow();
  serve();
}
