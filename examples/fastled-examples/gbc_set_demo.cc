#include <FastLED.h>
#define NUM_LEDS 75

CRGBArray<NUM_LEDS> leds;

void setup() { FastLED.addLeds<NEOPIXEL, 6>(leds, NUM_LEDS); }

void loop() {
  static uint8_t hue;
  for (int i = 0; i < NUM_LEDS / 2; i++) {
    leds.fadeToBlackBy(40);
    leds[i] = CHSV(hue++, 255, 255);
    leds(NUM_LEDS / 2, NUM_LEDS - 1) = leds(NUM_LEDS / 2 - 1, 0);
    FastLED.delay(33);
  }
}
