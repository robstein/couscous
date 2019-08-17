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
  WiFi.mode(WIFI_AP);
  WiFi.softAP("LEDLamp", "12345678");
  server.begin();

  FastLED.addLeds<LED_TYPE, DATA_PIN>(leds, NUM_LEDS)
      .setCorrection(TypicalLEDStrip);
  FastLED.setBrightness(BRIGHTNESS);
}

void loop() {
  Serial.begin(9600); // Start communication between the ESP8266-12E and the
                      // monitor window
  // IPAddress HTTPS_ServerIP = WiFi.softAPIP(); // Obtain the IP of the Server
  // Serial.print("Server IP is: ");             // Print the IP to the monitor
  // window Serial.println(HTTPS_ServerIP); delay(1000);

  WiFiClient client = server.available();
  if (!client) {
    return;
  }
  // Looking under the hood
  Serial.println("Somebody has connected :)");

  // Read what the browser has sent into a String class and print the request to
  // the monitor
  // String request = client.readStringUntil('\r');
  // char str[4];
  // strncpy(str, request[5], 3);
  // str[4] = '\0';
  // int i = String(str).indexOf(" ");
  // int val = 0;
  // if (i == -1) {
  //   val = atoi(str);
  // } else {
  //   char sstr[4 - i];
  //   strncpy(sstr, str, 3 - i);
  //   sstr[4 - i] = '\0';
  //   val = atoi(sstr);
  // }

  // leds[val] = CRGB(255, 0, 0);

  // Looking under the hood
  // Serial.println(request);

  client.flush();                      // clear previous info in the stream
  client.print("HTTP/1.1 200 OK\r\n"); // Send the response to the client
  delay(1);
  Serial.println("Client disonnected"); // Looking under the hood
}