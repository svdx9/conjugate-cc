export const checkAnswer = (userAns: string, expected: string): boolean => {
  return userAns.trim().toLowerCase() === expected.trim().toLowerCase();
};
