package io.cecg.bootcamp.initialisation.module_manager;

import io.cecg.bootcamp.initialisation.exception.ModuleManagerException;

import java.io.File;
import java.util.Collections;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class ModuleManager {
    public static Set<File> getAvailableModules(String moduleLocation) {
        File[] moduleFiles = new File(moduleLocation).listFiles();
        if (moduleFiles != null) {
            return Stream.of(moduleFiles)
                    .filter(File::isDirectory)
                    .collect(Collectors.toSet());
        } else {
            System.err.println("\nNo modules found in directory: " + moduleLocation + "\n");
            System.exit(1);
        }
        return Collections.emptySet();
    }

    public static void ensureModuleIsAvailable(Set<File> availableModules, Set<String> providedModuleNames) {
        providedModuleNames.forEach(providedModuleName -> {
            Set<String> availableModuleNames = availableModules.stream().map(File::getName).collect(Collectors.toSet());
            if (!availableModuleNames.contains(providedModuleName)) {
                throw new ModuleManagerException("Module " + providedModuleName + " does not exist. Available Modules: " + availableModuleNames);
            }
        });
    }
}
