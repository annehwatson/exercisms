// @ts-check

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let totalBirdCount = 0;
  for (let i = 0; i < birdsPerDay.length; i++) {
    totalBirdCount += birdsPerDay[i];
  }
  return totalBirdCount;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  const DAYS_IN_WEEK = 7
  let birdsInWeek = 0;
  let weekStartIndex = (week * DAYS_IN_WEEK) - DAYS_IN_WEEK;
  for (let i = weekStartIndex; i < weekStartIndex + DAYS_IN_WEEK; i++) {
    birdsInWeek += birdsPerDay[i];
  }
  return birdsInWeek;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let i = 0; i < birdsPerDay.length; i++) {
    if (i % 2 == 0) {
      birdsPerDay[i]++;
    }
  }
  return birdsPerDay;
}
