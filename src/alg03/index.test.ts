import { convertCelsiusToFahrenheit } from ".";

describe("convertCelsiusToFahrenheit", () => {
  it("convert 36 graus C return 96.8 graus F", () => {
    expect(convertCelsiusToFahrenheit(36)).toBe(96.8);
  });

  it("convert 15.5 graus C return 59.9 graus F", () => {
    expect(convertCelsiusToFahrenheit(15.5)).toBe(59.9);
  });

  it("convert -10 graus C return 14 graus F", () => {
    expect(convertCelsiusToFahrenheit(-10)).toBe(14);
  });
});
