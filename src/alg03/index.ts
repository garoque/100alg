export const convertCelsiusToFahrenheit = (celsius: number) => {
  return (celsius * 9) / 5 + 32;
};

console.log(convertCelsiusToFahrenheit(36));
