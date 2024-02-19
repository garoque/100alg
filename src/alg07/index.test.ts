import { calculateAreaSquare } from ".";

describe("calculateAreaSquare", () => {
  it("calculate area of 2", () => {
    expect(calculateAreaSquare(2)).toBe(4);
  });

  it("calculate negative area of 2", () => {
    expect(calculateAreaSquare(-2)).toBe(4);
  });
});
