import { calculateAverage } from ".";

describe("calculateAverage", () => {
  it("return the average of 2 plus 4 plus 6 and return 4", () => {
    expect(calculateAverage([2, 4, 6])).toBe(4);
    expect(calculateAverage([2, 4, 6])).not.toBe(4.6);
  });
});
