package io.cecg.initialisation.module_manager;

import io.cecg.initialisation.exception.ModuleManagerException;
import org.assertj.core.api.Assertions;
import org.junit.Test;

import java.io.File;
import java.net.URL;
import java.util.Set;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertEquals;

public class ModuleManagerTest {

    @Test
    public void shouldReturnAvailableModules() {
        // given
        URL testModule1 = getClass().getClassLoader().getResource("test_module");
        String path = testModule1.getPath();

        // when
        Set<File> testModule = ModuleManager.getAvailableModules(path);

        // then
        assertEquals(testModule.size(), 2);
    }

    @Test
    public void shouldNotThrowException_givenModuleIsAvailable() {
        // given
        Set<File> availableModuleFiles = Set.of(new File("module-1"), new File("module-2"), new File("module-3"));
        Set<String> providedModuleNames = Set.of("module-1", "module-2", "module-3");

        // when
        Throwable exception = Assertions.catchThrowable( () ->
                ModuleManager.ensureModuleIsAvailable(availableModuleFiles, providedModuleNames));

        // then
        assertThat(exception).isNull();
    }

    @Test
    public void shouldThrowException_givenUnavailableModule() {
        // given
        Set<File> availableModuleFiles = Set.of(new File("module-1"), new File("module-3"));
        Set<String> providedModuleNames = Set.of("module-1", "module-2", "module-3");

        // when
        Throwable exception = Assertions.catchThrowable( () ->
                ModuleManager.ensureModuleIsAvailable(availableModuleFiles, providedModuleNames));

        // then
        assertThat(exception).isNotNull()
                .isInstanceOf(ModuleManagerException.class)
                .hasMessageContaining("Module module-2 does not exist");
    }
}