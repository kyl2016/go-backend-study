ssh bp-global "curl -XGET -u elastic:bytepower -H 'Content-type: application/json' http://es-cluster.bytepower.app:9200/_cat/indices\?format\=json\&bytes='mb'\&h='index,store.size,pri'"
