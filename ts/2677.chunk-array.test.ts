import { describe, expect, test } from 'bun:test';
import { chunk } from './2677.chunk-array';

describe('chunk', () => {
  test.each([
    // [ arr, size, want]
    [[1, 2, 3], 3, [[1, 2, 3]]],
    [[1, 2, 3], 2, [[1, 2], [3]]],
    [[1, 2, 3], 1, [[1], [2], [3]]],
    [[1, 2, 3], 4, [[1, 2, 3]]],
    [[], 3, []],
  ])('chunk(%p, %p) to be %p', (arr, size, want) => {
    expect(chunk(arr, size)).toEqual(want);
  });
});
