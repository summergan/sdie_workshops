// @ts-check
import {test, expect} from '@playwright/test';

test('Search repository tags by keyword', async ({page}) => {
  const response = await page.goto('/user2/repo1/tags');
  await expect(response?.status()).toBe(200);

  await page.locator('input[name="q"]').fill('DELETE');
  await page.locator('input[name="q"]').press('Enter');

  await expect(page).toHaveURL(/\/user2\/repo1\/tags\?q=DELETE$/);
  await expect(page.locator('input[name="q"]')).toHaveValue('DELETE');
  await expect(page.locator('.tag-list-row-link')).toHaveText(['delete-tag']);
});
