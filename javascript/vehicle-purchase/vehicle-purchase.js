// @ts-check

/**
 * Determines whether or not you need a license to operate a certain kind of vehicle.
 *
 * @param {string} kind
 * @returns {boolean} whether a license is required
 */
export function needsLicense(kind) {
  return (kind == 'car' || kind == 'truck');
}

/**
 * Helps choosing between two options by recommending the one that
 * comes first in dictionary order.
 *
 * @param {string} option1
 * @param {string} option2
 * @returns {string} a sentence of advice which option to choose
 */
export function chooseVehicle(option1, option2) {
  let comparison = option1.localeCompare(option2);
  let chosenOption = "";
  if (comparison <= 0) {
    chosenOption = option1;
  } else {
    chosenOption = option2;
  }
  return chosenOption + " is clearly the better choice."
}

/**
 * Calculates an estimate for the price of a used vehicle in the dealership
 * based on the original price and the age of the vehicle.
 *
 * @param {number} originalPrice
 * @param {number} age
 * @returns {number} expected resell price in the dealership
 */
export function calculateResellPrice(originalPrice, age) {
  let remainingValuePercentage = 0.0;
  if (age < 3) {
    remainingValuePercentage = .8;
  } else if (age <= 10) {
    remainingValuePercentage = .7;
  } else {
    remainingValuePercentage = .5;
  }
  return remainingValuePercentage * originalPrice;
}
