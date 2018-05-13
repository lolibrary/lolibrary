import debounce from 'lodash/debounce';

export default class {
    constructor(element) {
        this.state = {
            categories: [],
            brands: [],
            features: [],
            years: [],
            colors: [],
            tags: [],
            search: '',
        };

        this.element = element;
    }

    async update(bounce) {
        const func = async () => {
            await Promise.all([
                async () => this.find('categories'),
                async () => this.find('brands'),
                async () => this.find('features'),
                async () => this.find('years'),
                async () => this.find('colors'),
                async () => this.find('tags'),
                async () => this.find('search'),
            ]);

            let event = new CustomEvent('SearchStateUpdated', {
                detail: this.state,
            });

            this.element.dispatchEvent(event);
        };

        if (bounce) {
            debounce(func, 250);
        } else {
            await func();
        }
    }

    find(id) {
        let element = document.getElementById(id);

        this.state[id] = $(element).val();
    }
}
