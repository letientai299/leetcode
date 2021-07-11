package leetcode.p1437;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SolutionTest {

  public static Stream<Arguments> getTests() {
    return Stream.of(
            Arguments.of(new int[]{1, 0, 0, 1, 0, 1}, 2, false),
            Arguments.of(new int[]{1, 0, 0, 0, 1, 0, 0, 1}, 2, true)
    );
  }

  @ParameterizedTest
  @MethodSource("getTests")
  void kLengthApart(int[] nums, int k, boolean want) {
    boolean got = new Solution().kLengthApart(nums, k);
    assertEquals(got, want);
  }
}
