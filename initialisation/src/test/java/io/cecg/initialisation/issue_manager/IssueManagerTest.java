package io.cecg.initialisation.issue_manager;

import org.junit.Test;
import org.kohsuke.github.GHIssue;
import org.kohsuke.github.GHIssueState;
import org.kohsuke.github.GHRepository;
import org.mockito.Mockito;

import java.io.IOException;
import java.util.Arrays;
import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.internal.verification.VerificationModeFactory.times;

public class IssueManagerTest {

    @Test
    public void shouldReturnIssuesByTitle() throws IOException {
        // given
        GHRepository mockRepository = Mockito.mock(GHRepository.class);
        GHIssue issue1 = Mockito.mock(GHIssue.class);
        GHIssue issue2 = Mockito.mock(GHIssue.class);

        String issueTitle1 = "Issue 1";
        String issueTitle2 = "Issue 2";

        when(mockRepository.getIssues(GHIssueState.ALL)).thenReturn(Arrays.asList(issue1, issue2));

        when(issue1.getTitle()).thenReturn(issueTitle1);
        when(issue2.getTitle()).thenReturn(issueTitle2);

        // when
        Map<String, GHIssue> issuesByTitle = IssueManager.getIssuesByTitle(mockRepository);

        // then
        assertThat(issuesByTitle).containsKeys(issueTitle1, issueTitle2);
        assertThat(issuesByTitle.get(issueTitle1)).isEqualTo(issue1);
        assertThat(issuesByTitle.get(issueTitle2)).isEqualTo(issue2);
        verify(mockRepository, times(1)).getIssues(GHIssueState.ALL);
    }
}