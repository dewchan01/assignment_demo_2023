import { check, group } from 'k6';
import http from 'k6/http';

export let options = {
  stages: [
    { duration: '1m', target: 20 },  // Ramp up to 20 virtual users over 1 minute
    { duration: '3m', target: 20 },  // Stay at 20 virtual users for 3 minutes
    { duration: '1m', target: 0 },   // Ramp down to 0 virtual users over 1 minute
  ],
};

export default function () {
  group('upload conversation to database', function () {
    const payload = JSON.stringify({
        chat: "jack:marcus",
        text: "hello marcus, i'm jack!",
        sender: "jack"
    });
    const headers = { 'Content-Type': 'application/json' };
    // locally test api with running http-service container
    // const res = http.post('http://localhost:8080/api/send', payload, { headers });
    const res = http.post('http://http-server:8080/api/send', payload, { headers });

    check(res, {
      'status is 200': (r) => r.status === 200
    });
  });

  group('pull message from database', function () {

    const payload = JSON.stringify({
      chat: "jack:marcus",
      cursor: 0,
      limit: 10,
      reverse: true
    });
    const headers = { 'Content-Type': 'application/json' };
    // locally test api with running http-service container
    // use http.request() instead of http.get() due to lack of "request payload params" in http.get()
    // const res = http.request('GET', 'http://localhost:8080/api/pull', payload, { headers });
    const res = http.request('GET', 'http://http-server:8080/api/pull', payload, { headers });
    check(res, {
      'status is 200': (r) => r.status === 200
    });
  })
    
}

