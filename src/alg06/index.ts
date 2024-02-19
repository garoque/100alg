export const conventMetersToCentimeters = (meter: number) => {
  const nonNegativeMeter = meter < 0 ? -1 * meter : meter;

  return nonNegativeMeter * 100;
};

console.log(conventMetersToCentimeters(-11));
