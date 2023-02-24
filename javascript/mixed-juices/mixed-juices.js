 const Juice = Object.freeze({
  Strawberry: "Pure Strawberry Joy",
  Energy: "Energizer",
  Green: "Green Garden",
  Tropical: "Tropical Island",
  Everything: "All or Nothing",
 });

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  if (!name) {
    throw new Error('Invalid juice name value');
  }

  switch (name) {
    case Juice.Strawberry:
      return 0.5;
    case Juice.Energy:
    case Juice.Green:
      return 1.5;
    case Juice.Tropical:
      return 3;
    case Juice.Everything:
      return 5;
    default:
      return 2.5;
  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  if (wedgesNeeded === 0 || limes.length === 0) {
    return 0;
  }

  let wedgesCut = 0;
  let limesCut = 0;
  for (let i = 0; i < limes.length; i++) {
    let currentLime = limes[i];
    switch (currentLime) {
      case 'small':
        wedgesCut += 6;
        break;
      case 'medium':
        wedgesCut += 8;
        break;
      case 'large':
        wedgesCut += 10;
        break;
      default:
       // no op
    }
    limesCut++;
    if (wedgesCut >= wedgesNeeded) {
      return limesCut;
    }
  }
  return limesCut;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  let timeLeftInShift = timeLeft;
  let remainingOrders = [...orders];
  for (let i = 0; i < orders.length; i++) {
    if (timeLeftInShift > 0) {
      timeLeftInShift -= timeToMixJuice(orders[i]);
      remainingOrders.shift();
    } else {
      break;
    }
  }
  return remainingOrders;
}
