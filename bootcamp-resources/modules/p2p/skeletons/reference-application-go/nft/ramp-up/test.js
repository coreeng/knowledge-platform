import http from 'k6/http';
import {check, sleep} from 'k6';
import {textSummary} from 'https://jslib.k6.io/k6-summary/0.0.2/index.js';


export const options = {
    stages: [
        { duration: '20s', target: 200 },
        { duration: '20s', target: 100 },
        { duration: '2m', target: 50 },
    ],
    thresholds: {
        'http_req_duration{status:200}': ['max>=0'],
        'http_req_duration{status:403}': ['max>=0'],
        'http_req_duration{status:404}': ['max>=0'],
        'http_req_duration{status:503}': ['max>=0'],
    },
};

const SERVICE_ENDPOINT = __ENV.SERVICE_ENDPOINT || 'http://service:8080';
const PUSH_GATEWAY_URL = __ENV.PUSH_GATEWAY_URL

export default function () {
    const res = http.get(`${SERVICE_ENDPOINT}/hello`);
    check(res, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}

export function handleSummary(data) {
    const timestamp = new Date().toISOString();
    const vusMax = data.metrics.vus_max.values.max;
    const status = data.root_group.checks[0].fails === 0 ? 'success' : 'failure';
    const httpReqDurationP95 = data.metrics.http_req_duration.values['p(95)']

    const metricString = `# TYPE k6_test_data counter\n` +
        `# HELP Counter to track k6 test run data\n` +
        `k6_test_data{status="${status}", vus_max="${vusMax}", p95_response_time="${httpReqDurationP95}", timestamp="${timestamp}"} 1\n`;

    http.post(PUSH_GATEWAY_URL + "/metrics/job/acceptance-criteria-nft-and-obs-k6-test-result", metricString, {
        headers: {"Content-Type": "text/plain"},
    });

    return {
        stdout: textSummary(data, {indent: '→', enableColors: true}),
    };
}