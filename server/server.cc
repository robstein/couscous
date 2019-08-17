#include "Arduino.h"
#include "FastLED.h"
#include <ESP8266WiFi.h>
#include <string.h>

WiFiServer server(80);

#define DATA_PIN 6
#define NUM_LEDS 760
CRGBArray<NUM_LEDS> leds;
#define BRIGHTNESS 96
#define LED_TYPE NEOPIXEL

void setup() {
  Serial.begin(9600);

  WiFi.mode(WIFI_AP);
  WiFi.softAP("LEDLamp", "12345678");
  server.begin();

  FastLED.addLeds<LED_TYPE, DATA_PIN>(leds, NUM_LEDS)
      .setCorrection(TypicalLEDStrip);
  FastLED.setBrightness(BRIGHTNESS);
}

void loop() {

  CRGB value = CRGB(0, 0, 0);
  int index = -1;

  WiFiClient client = server.available();
  // wait for a client (web browser) to connect
  if (client) {
    Serial.println("\n[Client connected]");
    while (client.connected()) {
      // read line by line what the client (web browser) is requesting
      if (client.available()) {
        String line = client.readStringUntil('\r');
        Serial.print(line);
        if (line.startsWith("GET /on/")) {
          String strval = line.substring(8, line.lastIndexOf(" "));
          Serial.println("[Debug] Received message to turn LED + " + strval +
                         " ON");
          index = strval.toInt();
          value = CRGB(255, 0, 0);
        } else if (line.startsWith("GET /off/")) {
          String strval = line.substring(9, line.lastIndexOf(" "));
          Serial.println("[Debug] Received message to turn LED + " + strval +
                         " OFF");
          index = strval.toInt();
        }
        // wait for end of client's request, that is marked with an empty line
        if (line.length() == 1 && line[0] == '\n') {
          client.print("HTTP/1.1 200 OK\r\nContent-Type: "
                       "text/html\r\nConnection: close\r\n\r\n");
          break;
        }
      }
    }
    delay(1); // give the web browser time to receive the data

    // close the connection:
    client.stop();
    Serial.println("[Client disonnected]");
  }

  if (index != -1) {
    leds[index] = value;
    FastLED.show();
    delay(1000);
  }
}
