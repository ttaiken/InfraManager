        $(function(){  
          $('#getservers').click(function(){  

            $.ajax({  

                url:'/infra/getservers',  

                type:'post',  

                success:function(data){  

                    var item;  
                    //$("#listserver tr:not(:first)").html("");
                    $("#listserver").bootstrapTable('destroy'); 
                    $("#listserver").show(); 
                    $.each(data,function(i,result){  

                        item = "<tr><td></td><td>"+result['hostname']+"</td><td>"+result['ip']+"</td><td>"+result['os']+"</td><td>"+result['platform']+"</td><td>"+"操作</td></tr>";  

                        $('#listserver').append(item);  

                    });  
                }  
            })  
          })
        })