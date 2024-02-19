import { calculateVolumeSphere } from ".";

describe("calculateVolumeSphere", () => {
  it("calculate spahre radius of 2.5", () => {
    expect(calculateVolumeSphere(2.5)).toBe(65.45);
  });
});
