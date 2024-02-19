import { calculateTrianguleArea } from ".";

describe("calculateTrianguleArea", () => {
  it("returns 10 with this perimeter", () => {
    expect(calculateTrianguleArea(10, 2)).toBe(10);
  });

  it("returns 6 with negative value", () => {
    expect(calculateTrianguleArea(-2, 6)).toBe(6);
    expect(calculateTrianguleArea(-2, 6)).not.toBe(-6);
  });

  it("returns 4.375 with float number", () => {
    expect(calculateTrianguleArea(2.5, 3.5)).toBe(4.375);
  });
});
