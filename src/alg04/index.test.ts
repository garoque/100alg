import { calculateRectanglePerimeter } from ".";

describe("calculateRectanglePerimeter", () => {
  it("returns 36 with this perimeter", () => {
    expect(calculateRectanglePerimeter(11, 7)).toBe(36);
  });

  it("returns 36 with negative value", () => {
    expect(calculateRectanglePerimeter(-11, 7)).toBe(36);
  });
});
