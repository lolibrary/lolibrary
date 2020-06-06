const axios = window.axios;
const endpoint = window.endpoint || 'https://lolibrary.org/api/v1/search';

export default async payload => {
    try {
        // payload should be an object to be POSTed to the search endpoint.

        return await axios.post(this.endpoint, payload);
    } catch (error) {
        console.error(error);
    }
};
