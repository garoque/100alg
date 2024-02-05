export const calculateAverage = (numbers: number[]) => {
  return (
    numbers.reduce((acc, currentValue) => {
      return acc + currentValue;
    }, 0) / numbers.length
  );
};
