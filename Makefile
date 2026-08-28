thrift=thrift
idl_path = idl/system-center

gen_service:
	kitex -module github.com/kouleen/system-center -service system.service  $(idl_path)/system.thrift
