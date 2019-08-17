#include "Arduino.h"
#include "FastLED.h"
#include <ESP8266WiFi.h>

#define DATA_PIN 6
#define NUM_LEDS 75

#define BUTTON_PIN 0

CRGB leds[NUM_LEDS];

void setup_wifi() {
  delay(10);
  // We start by connecting to a WiFi network
  Serial.println();
  Serial.print("Connecting to WiFi network");

  WiFi.mode(WIFI_STA);
  WiFi.begin("Robert's iPhone", "robfoo123");

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void setup() {
  Serial.begin(9600);
  pinMode(BUTTON_PIN, INPUT_PULLUP);
  FastLED.addLeds<NEOPIXEL, DATA_PIN>(leds, NUM_LEDS);

  setup_wifi();
}

int val = 0;
int prevVal = 0;
void onPress(int pin, void (*fn)()) {
  prevVal = val;
  val = digitalRead(pin);
  if ((val == 1) && (prevVal == 0)) {
    fn();
  }
}

int myIndex = 0;
void incrementIndex() { myIndex++; }

void loop() {
  onPress(BUTTON_PIN, incrementIndex);

  leds[myIndex] = CRGB(255, 0, 0);
  leds[myIndex + 1] = CRGB(0, 255, 0);
  leds[myIndex + 2] = CRGB(0, 0, 255);
  for (int i = 0; i < myIndex; i++) {
    leds[i] = CRGB(0, 0, 0);
  }
  for (int i = myIndex + 3; i < 75; i++) {
    leds[i] = CRGB(0, 0, 0);
  }
  FastLED.show();
  delay(10);
  Serial.print("Hello world.");
}
