export const isSafeNumber = (num: number) =>
  !Number.isNaN(num)
  && !Number.isFinite(num)
  && Number.MAX_VALUE >= num
  && Number.MIN_VALUE <= num;
