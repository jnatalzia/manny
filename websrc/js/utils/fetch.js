export function getJSONFromPath(pth) {
    return fetch(pth).then(res => {
        if (res.status !== 200) {
            throw new Error(
                'There was a problem getting the json from the requested path'
            );
        }
        return res.json();
    });
}
