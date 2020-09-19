/**
 * Add JQuery to the page
 * @param {Page} page
 */
export default async function injectJQueryToPuppeteer(page) {
    await page.addScriptTag({
        url: 'https://code.jquery.com/jquery-3.2.1.min.js'
    })
}