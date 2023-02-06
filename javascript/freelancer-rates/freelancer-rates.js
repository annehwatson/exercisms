// @ts-check

const DAILY_WORK_HOURS = 8;
const BILLABLE_DAYS_PER_MONTH = 22;

/**
 * The day rate, given a rate per hour
 *
 * @param {number} ratePerHour
 * @returns {number} the rate per day
 */
export function dayRate(ratePerHour) {
  return DAILY_WORK_HOURS * ratePerHour;
}

/**
 * Calculates the number of days in a budget, rounded down
 *
 * @param {number} budget: the total budget
 * @param {number} ratePerHour: the rate per hour
 * @returns {number} the number of days
 */
export function daysInBudget(budget, ratePerHour) {
  return Math.floor((budget/ratePerHour) / DAILY_WORK_HOURS);
}

/**
 * Calculates the discounted rate for large projects, rounded up
 *
 * @param {number} ratePerHour
 * @param {number} numDays: number of days the project spans
 * @param {number} discount: for example 20% written as 0.2
 * @returns {number} the rounded up discounted rate
 */
export function priceWithMonthlyDiscount(ratePerHour, numDays, discount) {
  let months = Math.floor(numDays / BILLABLE_DAYS_PER_MONTH);
  let fullPriceDays = numDays - (BILLABLE_DAYS_PER_MONTH * months);
  let discountHourlyRate = ratePerHour - (ratePerHour * discount);
  let discountedSubtotal = discountHourlyRate * DAILY_WORK_HOURS * BILLABLE_DAYS_PER_MONTH * months;
  let fullPriceSubtotal = fullPriceDays * DAILY_WORK_HOURS * ratePerHour;
  return Math.ceil(discountedSubtotal + fullPriceSubtotal);
}
