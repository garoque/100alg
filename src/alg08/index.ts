export const calculateVolumeSphere = (radius: number) => {
  const nonNegativeRadius = radius < 0 ? -1 * radius : radius;
  const volume = ((4 * Math.PI * Math.pow(nonNegativeRadius, 3)) / 3).toFixed(
    2
  );

  return Number(volume);
};

console.log(calculateVolumeSphere(6));
