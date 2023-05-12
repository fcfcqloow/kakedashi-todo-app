export const getEndDay = (year: number, month: number): number => {
  return new Date(year, month,  0).getDate();
};

export  const getStartWeek = (year: number, month: number): number => {
  return new Date(year, month - 1, 1).getDay();
};

export const nextMonth = (date: Date): Date => {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const afterYear = month + 1 === 13 ? year + 1 : year;
  const afterMonth = month + 1 === 13 ? 1 : month + 1;
  const result = new Date(date.getTime());
  result.setFullYear(afterYear);
  result.setMonth(afterMonth - 1);
  return result;
};

export const backMonth = (date: Date): Date => {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const beforeYear = month - 1 === 0 ? year - 1 : year;
  const beforeMonth = month - 1 === 0 ? 12 :  month - 1;
  const result = new Date(date.getTime());
  result.setFullYear(beforeYear);
  result.setMonth(beforeMonth - 1);
  return result;
};