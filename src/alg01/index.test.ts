import { addTwoNumbers } from ".";

describe("addTwoNumbers", () => {
  it("returns 4 with result", () => {
    expect(addTwoNumbers(2, 2)).toBe(4);
  });

  it("returns 5 with result", () => {
    expect(addTwoNumbers(15, -10)).toBe(5);
  });

  it("returns false for 2 + 2 equals 7", () => {
    expect(addTwoNumbers(2, 2)).not.toBe(7);
  });
});
