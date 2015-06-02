var user = {
    fName: 'John',
    lName: 'Doe',
    links: {
        blog: 'www.blogger.com',
        facebook: 'www.facebook.com'
    }
};

for (var i in user.links) {
    console.log(i, ' - ', user.links[i]);
}