import injectJQueryToPuppeteer from "./actions/inject-jquery";
import puppeteer from "puppeteer";

/**
 * Crawler-Facade of Horizon
 */
const Crawler = {}

/**
 * Run the crawler
 * @param {String} url
 * @param {Object} rules
 * @return {Object|null}
 */
Crawler.make = async function (
    url,
    rules
) {
    const resultBucket = {
        status: true,
        // data will stay below here
    }

    try {
        const [page, browser] = puppeteerCall(url)

        // injectJquery
        if (!rules.injectJQuery || rules.injectJQuery === true) {
            await injectJQueryToPuppeteer(page)
        }

        /**
         * Loop through the actions and handle it
         * @var {Object} actionRule
         */
        for (const actionRule of rules.actions) {

        }


        // when everything is done, close the browser and result the result
        browser.close()
        return resultBucket
    } catch (e) {
        // TODO: A logger should be added here
        resultBucket.status = false
        resultBucket.errorMessage = e.message
        return resultBucket
    }

}

/**
 * Call Puppeteer
 * @param url
 * @param options
 * @returns {Promise<(Page|Browser)[]>}
 */
async function puppeteerCall(url, options = {}) {
    const browser = await puppeteer.launch(options)

    // create page then go to it
    const page = await browser.newPage();
    const httpResponse = await page.goto(url, { waitUntil: 'networkidle2' });

    if (!httpResponse.ok()) {
        await browser.close()
        throw new Error('Failed to continue due to the response is not ok. Status: ' + httpResponse.status())
    }

    return [page, browser];
}

export {
    Crawler
}