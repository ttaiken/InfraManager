$(function(){
    var $table = $('#listserver');
    $table.bootstrapTable({
    url: "/infra/getservers", 
    method: 'get',  
    dataType: "json",
    pagination: true, //分页
    //queryParams: queryParams,  //{limit, offset, search, sort, order}
    //queryParamsType: '',
    search: true, //显示搜索框
    sidePagination: "server", //服务端处理分页
    pageList: [10, 25, 50, 100],
    sortable: true, 
    showToggle: false,
    
    columns: [
      {
        title: 'hostname',
          field: 'hostname',
          align: 'center',
          valign: 'middle',
          sortable: true,
          visible: true,
          editable : true,
      }, 
      {
          title: 'ip',
          field: 'ip',
          align: 'center',
          valign: 'middle',
      }, 
      {
          title: 'os',
          field: 'os',
          align: 'center'
      },
      {
          title: 'platform',
          field: 'platform',
          align: 'center',
      },
      {
          title: '操作',
          field: 'hostname',
          align: 'center',
          formatter:function(value,row,index){  
            var e = '<a href="#" mce_href="#" onclick="edit(\''+ row.hostname +'\')">编辑</a> ';  
            var d = '<a href="#" mce_href="#" onclick="del(\''+ row.hostname +'\')">删除</a> ';  
            return e+d; 
          } 
      }
    ]
  });
})

function del(params) {
      window.alert(params)
};

function edit(params){
        window.alert(params)

};