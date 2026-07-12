// @ts-check
import {test, expect} from '@playwright/test';

async function loginUser(page, workerInfo, user) {
  const response = await page.goto('/user/login');
  await expect(response?.status()).toBe(200);

  await page.locator('input[name=user_name]').fill(user);
  await page.locator('input[name=password]').fill('password');
  await page.locator('form button.ui.primary.button:visible').click();
  await expect(page).toHaveURL(`${workerInfo.project.use.baseURL}/`);
}

test('Actions status filter menu shows localized labels and icons', async ({page}, workerInfo) => {
  await loginUser(page, workerInfo, 'user5');

  const settingsResponse = await page.goto('/user5/repo4/settings');
  await expect(settingsResponse?.status()).toBe(200);

  const advancedForm = page.locator('form:has(input[name="action"][value="advanced"])');
  await expect(advancedForm).toBeVisible();

  const actionsCheckbox = advancedForm.locator('input[name="enable_actions"]');
  if (!await actionsCheckbox.isChecked()) {
    await actionsCheckbox.check({force: true});
    await advancedForm.locator('button.ui.primary.button').click();
    await page.waitForLoadState('domcontentloaded');
    await expect(page.locator('.ui.positive.message.flash-success')).toContainText('The repository settings have been updated.');
  }

  const actionsResponse = await page.goto('/user5/repo4/actions');
  await expect(actionsResponse?.status()).toBe(200);

  const statusDropdown = page.locator('.ui.secondary.filter.menu .ui.dropdown.jump.item').last();
  await statusDropdown.click();
  await expect(statusDropdown.locator('.menu')).toBeVisible();

  for (const [status, label] of [
    ['1', 'Success'],
    ['2', 'Failure'],
    ['3', 'Canceled'],
    ['4', 'Skipped'],
    ['5', 'Waiting'],
    ['6', 'Running'],
    ['7', 'Blocked'],
  ]) {
    const item = statusDropdown.locator(`a.item[href*="status=${status}"]`);
    await expect(item).toHaveCount(1);
    await expect(item.locator('span[data-tooltip-content] svg')).toHaveCount(1);
    await expect(item.locator('span').last()).toHaveText(label);
  }
});
