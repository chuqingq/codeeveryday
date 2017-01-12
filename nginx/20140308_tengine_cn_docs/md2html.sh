docs="ngx_dso_module_cn ngx_http_concat_cn ngx_http_core_module_cn ngx_http_tfs_module_cn ngx_http_trim_filter_module_cn ngx_http_upstream_check_module_cn ngx_http_upstream_consistent_hash_module_cn ngx_http_upstream_session_sticky_module_cn ngx_procs_module_cn TFS_RESTful_API_cn"

for f in $docs; do
    echo "<html><head><meta charset=\"UTF-8\"></head>" > $f.html
    markdown $f.md >> $f.html
    echo "</html>" >> $f.html
done

