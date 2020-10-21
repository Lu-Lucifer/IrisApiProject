<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddUser">添加用户</el-button>

    <el-table :data="usersList" style="width: 100%;margin-top:30px;" border>
      <el-table-column align="center" label="ID" width="220">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="用户标识" width="220">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="用户名称" width="220">
        <template slot-scope="scope">
          {{ scope.row.username }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="描述">
        <template slot-scope="scope">
          {{ scope.row.introduction }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button type="primary" size="small" @click="handleEdit(scope)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'编辑用户':'添加用户'">
      <el-form :model="user" label-width="80px" label-position="left">
        <el-form-item label="标识">
          <el-input v-model="user.name" placeholder="用户标识" />
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="user.username" placeholder="用户名称" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="user.password" placeholder="密码" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input
            v-model="user.introduction"
            :autosize="{ minRows: 2, maxRows: 4}"
            type="textarea"
            placeholder="用户描述"
          />
        </el-form-item>
        <el-form-item label="角色">
          <el-tree
            ref="tree"
            :check-strictly="checkStrictly"
            :data="routesData"
            :props="defaultProps"
            show-checkbox
            node-key="id"
            class="permission-tree"
          />
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="confirmUser">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>

import { deepClone } from '@/utils'
import { getRoles, getUsers, addUser, deleteUser, updateUser } from '@/api/users'

const defaultUser = {
  id: 0,
  name: '',
  username: '',
  introduction: '',
  roles: [],
  role_ids: []
}

export default {
  data() {
    return {
      user: Object.assign({}, defaultUser),
      roles: [],
      usersList: [],
      dialogVisible: false,
      dialogType: 'new',
      checkStrictly: false,
      defaultProps: {
        children: 'children',
        label: 'title'
      }
    }
  },
  computed: {
    routesData() {
      return this.roles
    }
  },
  created() {
    // Mock: get all roles and users list from server
    this.getRoles()
    this.getUsers()
  },
  methods: {
    async getRoles() {
      const res = await getRoles()
      this.serviceRoles = res.data
      this.roles = this.generateRoles(res.data)
    },
    async getUsers() {
      const res = await getUsers()
      this.usersList = res.data
    },

    // Reshape the routes structure so that it looks the same as the sidebar
    generateRoles(roles) {
      const res = []
      if (roles == null || roles.length === 0) {
        return res
      }
      for (const perm of roles) {
        if (perm.hidden) { continue }
        const data = {
          id: perm.id,
          title: perm.name
        }
        res.push(data)
      }
      return res
    },
    generateArr(roles) {
      const data = []
      roles.forEach(perm => {
        data.push(perm)
      })
      return data
    },
    handleAddUser() {
      this.user = Object.assign({}, defaultUser)
      if (this.$refs.tree) {
        this.$refs.tree.setCheckedNodes([])
      }
      this.dialogType = 'new'
      this.dialogVisible = true
    },
    handleEdit(scope) {
      this.dialogType = 'edit'
      this.dialogVisible = true
      this.checkStrictly = true
      this.user = deepClone(scope.row)
      this.$nextTick(() => {
        const roles = this.generateRoles(this.user.roles)
        this.$refs.tree.setCheckedNodes(this.generateArr(roles))
        // set checked state of a node not affects its father and child nodes
        this.checkStrictly = false
      })
    },
    handleDelete({ $index, row }) {
      this.$confirm('确认移除这个用户吗?', 'Warning', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          await deleteUser(row.id)
          this.usersList.splice($index, 1)
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        .catch(err => { console.error(err) })
    },
    generateTree(roles, checkedKeys) {
      const res = []
      for (const perm of roles) {
        if (checkedKeys.includes(perm.id)) {
          res.push(perm.id)
        }
      }
      return res
    },
    async confirmUser() {
      const isEdit = this.dialogType === 'edit'

      const checkedKeys = this.$refs.tree.getCheckedKeys()
      this.user.role_ids = this.generateTree(deepClone(this.serviceRoles), checkedKeys)

      if (isEdit) {
        const { code, data } = await updateUser(this.user.id, this.user)
        // eslint-disable-next-line eqeqeq
        if (code === 200) {
          this.user = data
          for (let index = 0; index < this.usersList.length; index++) {
            if (this.usersList[index].id === this.user.id) {
              this.usersList.splice(index, 1, Object.assign({}, this.user))
              break
            }
          }
        }
      } else {
        const { code, data } = await addUser(this.user)
        if (code === 200) {
          this.user = data
          this.usersList.push(this.user)
        }
      }

      const { introduction, username, name } = this.user
      this.dialogVisible = false
      this.$notify({
        title: 'Success',
        dangerouslyUseHTMLString: true,
        message: `
            <div>用户标识: ${name}</div>
            <div>用户名称: ${username}</div>
            <div>描述: ${introduction}</div>
          `,
        type: 'success'
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.app-container {
  .users-table {
    margin-top: 30px;
  }
  .permission-tree {
    margin-bottom: 30px;
  }
}
</style>
