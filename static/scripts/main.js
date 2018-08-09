var articles = new Vue({
    el: '#articles',
    count: '',
    data: {
        posts: [],
    },
    mounted: function () {
        var a = axios.get('http://localhost:8080/1').then(function (res) {
            // console.log(res.data);
            articles.posts = res.data;
            // console.log(articles.posts);
        }).catch(function (error) {
            articles.posts = {"err": error.message}
        });
        // var a = axios.get('../dummy/main.json').then(function (res) {
        //     console.log(res.data);
        //     articles.posts = res.data;
        //     console.log(articles.posts);
        // });
    }
});
articles.count = 0;