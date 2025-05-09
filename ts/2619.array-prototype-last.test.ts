import { expect, test } from 'bun:test';
import './2619.array-prototype-last';

test('Array.last', () => {
  const arr = [1, 3, 2];
  expect(arr.last()).toBe(2);
});

test.each([
  { arr: [1, 2, 3], expected: 3 },
  { arr: [], expected: -1 },
])('Array.last(%j)', ({ arr, expected }) => {
  expect(arr.last()).toBe(expected);
});
