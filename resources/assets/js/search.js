const axios = window.axios;

export default class {
    constructor(element, endpoint) {
        this.element = element;
        this.endpoint = endpoint;

        this.element.registerEventListener('SearchStateUpdated', this.search);
    }

    async search(event) {
        const request = event.detail;

        try {
            const result = await axios.post(this.endpoint, request);

            await this.update(result);
        } catch (error) {
            console.error(error);
        }
    }

    async update(data) {
        console.log(data);
    }
}
