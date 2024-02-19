export const calculateTrianguleArea = (base: number, height: number) => {
  const nonNegativeBase = base < 0 ? -1 * base : base;
  const nonNegativeHeight = height < 0 ? -1 * height : height;

  return (nonNegativeBase * nonNegativeHeight) / 2;
};

console.log(calculateTrianguleArea(-11, -7));
