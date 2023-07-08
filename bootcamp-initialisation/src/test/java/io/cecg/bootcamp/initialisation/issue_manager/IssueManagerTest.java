package io.cecg.bootcamp.initialisation.issue_manager;

import io.cecg.bootcamp.initialisation.model.CmdArguments;
import org.junit.Test;
import org.kohsuke.github.GHIssue;
import org.kohsuke.github.GHIssueState;
import org.kohsuke.github.GHRepository;
import org.mockito.Mockito;

import java.io.File;
import java.io.IOException;
import java.util.*;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.*;
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

    @Test
    public void testCreateModuleLabels() throws IOException {
        // given
        CmdArguments cmdArguments = new CmdArguments.Builder().setModules(new HashSet<>(Arrays.asList("module-1", "module-2"))).build();
        GHRepository mockedGhRepository = Mockito.mock(GHRepository.class);

        File module1 = Mockito.mock(File.class);
        File module2 = Mockito.mock(File.class);

        Set<File> modules = Set.of(module1, module2);

        GHIssue module1Issue = Mockito.mock(GHIssue.class);
        GHIssue module2Issue = Mockito.mock(GHIssue.class);
        Mockito.when(module1Issue.getTitle()).thenReturn("module-1");
        Mockito.when(module2Issue.getTitle()).thenReturn("module-2");

        Map<String, GHIssue> issuesByTitle = Map.of("module-1", module1Issue, "module-2", module2Issue);

        // when
        IssueManager.createModuleIssues(cmdArguments, mockedGhRepository, modules, issuesByTitle);

        // then
        verify(mockedGhRepository, times(2)).createLabel(anyString(), anyString());
    }
}