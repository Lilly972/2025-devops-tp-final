import { test, expect } from '@playwright/test';

test.describe('Christmas Gift List App', () => {
  test('should display the home page with Christmas theme', async ({ page }) => {
    await page.goto('/');
    
    // Check for Christmas title
    await expect(page.locator('h1')).toContainText('Christmas Gift List');
    
    // Check for snowflakes
    await expect(page.locator('.snowflake').first()).toBeVisible();
  });

  test('should add a new person', async ({ page }) => {
    await page.goto('/');
    
    // Fill in the person name
    await page.fill('input[placeholder*="person\'s name"]', 'Test Person');
    
    // Click add button
    await page.click('button:has-text("Add Person")');
    
    // Verify person was added
    await expect(page.locator('text=Test Person')).toBeVisible();
  });

  test('should navigate to person\'s gift page and add a gift', async ({ page }) => {
    await page.goto('/');
    
    // Add a person first
    await page.fill('input[placeholder*="person\'s name"]', 'Gift Test Person');
    await page.click('button:has-text("Add Person")');
    
    // Click view gifts button
    await page.click('text=Gift Test Person >> .. >> button:has-text("View Gifts")');
    
    // Should be on the person's page
    await expect(page.locator('h1')).toContainText('Gift Ideas for Gift Test Person');
    
    // Add a gift idea
    await page.fill('input[placeholder*="Gift title"]', 'Test Gift');
    await page.fill('textarea[placeholder*="Gift description"]', 'This is a test gift description');
    await page.click('button:has-text("Add Gift Idea")');
    
    // Verify gift was added
    await expect(page.locator('text=Test Gift')).toBeVisible();
    await expect(page.locator('text=This is a test gift description')).toBeVisible();
  });

  test('should select a gift as final', async ({ page }) => {
    await page.goto('/');
    
    // Add a person
    await page.fill('input[placeholder*="person\'s name"]', 'Selection Test');
    await page.click('button:has-text("Add Person")');
    
    // Navigate to gifts page
    await page.click('text=Selection Test >> .. >> button:has-text("View Gifts")');
    
    // Add two gifts
    await page.fill('input[placeholder*="Gift title"]', 'Gift 1');
    await page.fill('textarea[placeholder*="Gift description"]', 'First gift');
    await page.click('button:has-text("Add Gift Idea")');
    
    await page.fill('input[placeholder*="Gift title"]', 'Gift 2');
    await page.fill('textarea[placeholder*="Gift description"]', 'Second gift');
    await page.click('button:has-text("Add Gift Idea")');
    
    // Select the second gift
    await page.locator('text=Gift 2 >> .. >> button:has-text("Select as Final Gift")').click();
    
    // Verify selection banner appears
    await expect(page.locator('text=Final Gift Selected!')).toBeVisible();
    await expect(page.locator('text=Gift 2')).toHaveCount(2); // Once in banner, once in card
  });

  test('should delete a person', async ({ page }) => {
    await page.goto('/');
    
    // Add a person
    await page.fill('input[placeholder*="person\'s name"]', 'Delete Me');
    await page.click('button:has-text("Add Person")');
    
    // Verify person exists
    await expect(page.locator('text=Delete Me')).toBeVisible();
    
    // Click delete button and confirm
    page.on('dialog', dialog => dialog.accept());
    await page.locator('text=Delete Me >> .. >> button[aria-label*="Delete"]').click();
    
    // Verify person was deleted
    await expect(page.locator('text=Delete Me')).not.toBeVisible();
  });
});
