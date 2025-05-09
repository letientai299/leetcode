import { describe, expect, test } from 'bun:test';
import { createCounter } from './2620.counter';

describe('createCounter', () => {
  test.each([
    // [ n ]
    [10],
  ])('%p', (n) => {
    let counter = createCounter(n);
    expect(counter()).toBe(n);
    expect(counter()).toBe(n + 1);
    expect(counter()).toBe(n + 2);
  });
});
