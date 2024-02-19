import { conventMetersToCentimeters } from ".";

describe("conventMetersToCentimeters", () => {
  it("calculate 3 meters to centimeters", () => {
    expect(conventMetersToCentimeters(3)).toBe(300);
  });

  it("calculate negative 3 meters to centimeters", () => {
    expect(conventMetersToCentimeters(-3)).toBe(300);
    expect(conventMetersToCentimeters(-3)).not.toBe(-300);
  });
});
