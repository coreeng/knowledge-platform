package io.cecg.initialisation.github_manager;

import io.cecg.initialisation.exception.GitHubManagerException;
import io.cecg.initialisation.model.CmdArguments;
import org.kohsuke.github.GHOrganization;
import org.kohsuke.github.GHRepository;
import org.kohsuke.github.GitHub;
import org.kohsuke.github.GitHubBuilder;

import java.io.IOException;

public class GitHubManager {
    public static GHRepository getOrCreateRepository(CmdArguments cmdArguments) {
        String token = cmdArguments.getGitToken();
        try {
            GitHub github = new GitHubBuilder().withOAuthToken(token).build();
            GHOrganization organization = github.getOrganization(cmdArguments.getOrg());
            return createIfNotPresent(cmdArguments, organization);
        } catch (IOException ioException) {
            throw new GitHubManagerException("Issue occurred while contacting Github", ioException);
        }
    }

    private static GHRepository createIfNotPresent(CmdArguments cmdArguments, GHOrganization organization) throws IOException {
        GHRepository ghRepository = organization.getRepository(cmdArguments.getBootcampeeRepo());

        if (ghRepository == null) {
            ghRepository = organization.createRepository(cmdArguments.getBootcampeeRepo())
                    .owner(cmdArguments.getOrg())
                    .private_(true)
                    .create();
        }
        return ghRepository;
    }
}
