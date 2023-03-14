const storage = {
    set: (key: string, value: unknown) => {
        localStorage.setItem(key, JSON.stringify(value));
    },
    get: (key: string) => {
        const value = localStorage.getItem(key);
        if (value === "undefined" || value === "null"){
            return null
        }
        return value ? JSON.parse(value) : null;
    },
    remove: (key: string) => {
        localStorage.removeItem(key);
    },
    clear: () => {
        localStorage.clear();
    },
};

export default storage;