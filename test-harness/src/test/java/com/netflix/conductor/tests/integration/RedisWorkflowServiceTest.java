/*
 * Copyright 2016 Netflix, Inc.
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package com.netflix.conductor.tests.integration;

import com.netflix.conductor.tests.utils.RedisTestRunner;
import org.junit.runner.RunWith;

import java.util.Map;

@RunWith(RedisTestRunner.class)
public class RedisWorkflowServiceTest extends AbstractWorkflowServiceTest {

    @Override
    String startOrLoadWorkflowExecution(String snapshotResourceName, String workflowName, int version, String correlationId, Map<String, Object> input, String event, Map<String, String> taskToDomain) {
        return workflowExecutor.startWorkflow(workflowName, version, correlationId, input, null, event, taskToDomain);
    }
}