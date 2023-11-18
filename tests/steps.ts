import { Page, expect } from '@playwright/test';
import { createBdd } from 'playwright-bdd';
import { test } from './fixtures';

const { Given, When, Then } = createBdd(test);

Given('an empty working directory', async ({ app }) => {
  await app.useTempDir();
  console.log("CWD:", app.cwd);
});

When('the user runs the app', async ({ app }) => {
  await app.run();
});

Then('debug wait {int}', async ({}, ms: number) => {
  await new Promise(resolve => setTimeout(resolve, ms));
});

Then('the app logs {string}', async ({ app }, log: string) => {
  await expect(async () => {
    expect(app.stderr).toContain(log);
  }).toPass();
});

Given('a running API', async ({ app }) => {
  await app.run();
  await expect(async () => {
    expect(app.stderr).toContain("api at :8080/");
  }).toPass();
});

When('the API goes down', async ({ app }) => {
  await app.stop();
});

When('the API comes back up', async ({ app }) => {
  await app.run();
});

When('the user waits for {int} seconds', async ({ page }, sec: number) => { 
  await page.waitForTimeout(sec * 1000);
});

When('the user opens the home page', async ({ app }) => {
  await app.goto("/");
});

Then('the page shows a progress bar', async ({ page }) => {
  await expect(page.locator("#content").getByRole('progressbar')).toBeVisible();
});

Then('the page shows {string}', async ({ page }, text) => {
  await expect(page.getByText(text)).toBeVisible();
});

Then('the page does not show {string}', async ({ page }, text: string) => {
  await expect(page.getByText(text)).not.toBeVisible();
});

When('the user switches away and back to the page', async ({ page }) => {
  await page.evaluate(() => {
    document.dispatchEvent(new Event('visibilitychange'))
  })
});

When('the user clicks {string}', async ({ page }, text: string) => {
  await page.getByText(text).click();
});

When('the user adds a folder {string}', async ({ app }, name: string) => {
  await app.addDir(name);
});

When('the user clicks "Retry', async ({ page }) => {
  await page.getByRole('button', { name: 'Retry' }).click();
});
  