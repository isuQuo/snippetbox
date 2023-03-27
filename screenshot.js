const puppeteer = require('puppeteer');
const waitOn = require('wait-on');

(async () => {
  await waitOn({
    resources: ['http://localhost:8080'],
    timeout: 30000, // 30 seconds
  });
  
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto('http://localhost:8080');
  await page.screenshot({path: 'screenshot.png'});
  await browser.close();
})();
