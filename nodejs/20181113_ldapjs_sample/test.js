var ldap = require('ldapjs');

console.log('before createClient');
var client = ldap.createClient({
  url: 'ldap://192.168.0.200:389'
});

console.log('before bind');
client.bind('cn=admin,dc=dilu,dc=com', '5tgb6yhn%TGB^YHN', function(err) {
  console.log('bind error:', err);

  var opts = {
    // filter: "(uid=songxufang24)", //查询条件过滤器，查找uid=kxh的用户节点
    scope: "sub", //查询范围
    // timeLimit: 500 //查询超时
  };
  client.search("dc=dilu,dc=com", opts, function(err, res) {
    console.log('search result: ', err, res);
    res.on('searchEntry', function(entry) {
      console.log('entry: ' + JSON.stringify(entry.object));
    });
    res.on('searchReference', function(referral) {
      console.log('referral: ' + referral.uris.join());
    });
    res.on('error', function(err) {
      console.error('error: ' + err.message);
    });
    res.on('end', function(result) {
      console.log('status: ' + result.status);
    });
  });
});

