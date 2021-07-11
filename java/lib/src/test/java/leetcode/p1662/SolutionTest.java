package leetcode.p1662;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SolutionTest {

  public static Stream<Arguments> getTests() {
    return Stream.of(
            Arguments.of(
                    new String[]{"ab", "c"},
                    new String[]{"a", "bc"},
                    true
            ),

            Arguments.of(
                    new String[]{"ab", "c"},
                    new String[]{"a", "bcd"},
                    false
            )
    );
  }

  @ParameterizedTest
  @MethodSource("getTests")
  void arrayStringsAreEqual(String[] word1, String[] word2, boolean want) {
    boolean got = new Solution().arrayStringsAreEqual(word1, word2);
    assertEquals(got, want);
  }
}
