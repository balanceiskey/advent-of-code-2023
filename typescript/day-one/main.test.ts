import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { getScore, getScoreCorrected } from "./main.ts";

Deno.test("getScore returns correct value in happy path", () => {
  const input = "1abc2";
  const expected = 12;

  assertEquals(expected, getScore(input));
});

Deno.test("getScore returns correct value with more than 2 digits", () => {
  const input = "a1b2c3d4e5f";
  const expected = 15;

  assertEquals(expected, getScore(input));
});

Deno.test(
  "getScore returns correct value when only a single digit exists",
  () => {
    const input = "treb7uchet";
    const expected = 77;

    assertEquals(expected, getScore(input));
  }
);

Deno.test("getScoreCorrected returns correct value in variety of cases", () => {
  const testCases: Record<string, number> = {
    two1nine: 29,
    eightwothree: 83,
    abcone2threexyz: 13,
    xtwone3four: 24,
    "4nineeightseven2": 42,
    zoneight234: 14,
    "7pqrstsixteen": 76,
  };

  Object.keys(testCases).map((input) => {
    const expected = testCases[input];

    assertEquals(expected, getScoreCorrected(input));
  });
});

Deno.test("getScoreCorrected manages overlaps", () => {
  const input = "1twoneone";
  const expected = 11;
  assertEquals(expected, getScoreCorrected(input));
});
