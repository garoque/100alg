export const calculateRectanglePerimeter = (base: number, height: number) => {
  const nonNegativeBase = base < 0 ? -1 * base : base;
  const nonNegativeHeight = height < 0 ? -1 * height : height;

  return 2 * (nonNegativeBase + nonNegativeHeight);
};

console.log(calculateRectanglePerimeter(-11, -7));
