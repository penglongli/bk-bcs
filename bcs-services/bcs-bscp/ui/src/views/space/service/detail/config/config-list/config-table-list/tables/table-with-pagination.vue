<script lang="ts" setup>
  import { ref, watch, onMounted } from 'vue'
  import { storeToRefs } from 'pinia'
  import { InfoBox } from "bkui-vue/lib";
  import { useConfigStore } from '../../../../../../../../store/config'
  import { IConfigItem, IConfigListQueryParams } from '../../../../../../../../../types/config'
  import { getConfigList, deleteServiceConfigItem } from '../../../../../../../../api/config'
  import { getConfigTypeName } from '../../../../../../../../utils/config'
  import { datetimeFormat } from '../../../../../../../../utils/index'
  import StatusTag from './status-tag';
  import EditConfig from '../edit-config.vue'
  import VersionDiff from '../../../components/version-diff/index.vue'

  const configStore = useConfigStore()
  const { versionData } = storeToRefs(configStore)

  const props = defineProps<{
    bkBizId: string;
    appId: number;
    searchStr: string;
  }>()

  const loading = ref(false)
  const configList = ref<IConfigItem[]>([])
  const editPanelShow = ref(false)
  const activeConfig = ref(0)
  const isDiffPanelShow = ref(false)
  const diffConfig = ref(0)
  const pagination = ref({
    current: 1,
    count: 0,
    limit: 10,
  })

  watch(() => versionData.value.id, () => {
    getListData()
  })

  watch(() => props.searchStr, () => {
    refresh()
  })

  onMounted(() => {
    getListData()
  })

  const getListData = async () => {
    loading.value = true
    try {
      const params: IConfigListQueryParams = {
        start: (pagination.value.current - 1) * pagination.value.limit,
        limit: pagination.value.limit
      }
      if (props.searchStr) {
        params.searchKey = props.searchStr
      }
      if (versionData.value.id !== 0) {
        params.release_id = versionData.value.id
      }
      const res = await getConfigList(props.bkBizId, props.appId, params)
      configList.value = res.details
      pagination.value.count = res.count
    } catch (e) {
      console.error(e)
    } finally {
      loading.value = false
    }
  }

  const handleEdit = (config: IConfigItem) => {
    activeConfig.value = config.id
    editPanelShow.value = true
  }

  const handleDiff = (config: IConfigItem) => {
    diffConfig.value = config.id
    isDiffPanelShow.value = true
  }

  const handleDel = (config: IConfigItem) => {
    InfoBox({
      title: `确认是否删除配置项【${config.spec.name}】?`,
      infoType: "danger",
      headerAlign: "center" as const,
      footerAlign: "center" as const,
      onConfirm: async () => {
        await deleteServiceConfigItem(config.id, props.bkBizId, props.appId)
        if (configList.value.length === 1 && pagination.value.current > 1) {
          pagination.value.current -= 1
        }
        getListData();
      },
    } as any);
  }

  const handlePageLimitChange = (limit: number) => {
    pagination.value.limit = limit
    refresh()
  }

  const refresh = (current: number = 1) => {
    pagination.value.current = current
    getListData()
  }

  defineExpose({
    refresh
  })

</script>
<template>
  <bk-loading :loading="loading">
    <bk-table
      v-if="!loading"
      :border="['outer']"
      :data="configList"
      :pagination="pagination"
      @page-limit-change="handlePageLimitChange"
      @page-change="refresh($event)">
      <bk-table-column label="配置文件名" prop="spec.name" :sort="true" :min-width="240" show-overflow-tooltip>
        <template #default="{ row }">
          <bk-button
            v-if="row.spec"
            text
            theme="primary"
            :disabled="row.file_state === 'DELETE'"
            @click="handleEdit(row)">
            {{ row.spec.name }}
          </bk-button>
        </template>
      </bk-table-column>
      <bk-table-column label="配置文件路径" prop="spec.path" show-overflow-tooltip></bk-table-column>
      <bk-table-column label="配置文件格式">
        <template #default="{ row }">
          {{ getConfigTypeName(row.spec?.file_type) }}
        </template>
      </bk-table-column>
      <bk-table-column label="创建人" prop="revision.creator"></bk-table-column>
      <bk-table-column label="修改人" prop="revision.reviser"></bk-table-column>
      <bk-table-column label="修改时间" :sort="true" :width="220">
        <template #default="{ row }">
          <span v-if="row.revision">{{ datetimeFormat(row.revision.update_at) }}</span>
        </template>
      </bk-table-column>
      <bk-table-column v-if="versionData.id === 0" label="变更状态">
        <template #default="{ row }">
          <StatusTag :status="row.file_state" />
        </template>
      </bk-table-column>
      <bk-table-column label="操作" fixed="right">
        <template #default="{ row }">
          <div class="operate-action-btns">
            <bk-button :disabled="row.file_state === 'DELETE'" text theme="primary" @click="handleEdit(row)">{{ versionData.id === 0 ? '编辑' : '查看' }}</bk-button>
            <bk-button v-if="versionData.status.publish_status !== 'editing'" text theme="primary" @click="handleDiff(row)">对比</bk-button>
            <bk-button v-if="versionData.id === 0" text theme="primary" :disabled="row.file_state === 'DELETE'" @click="handleDel(row)">删除</bk-button>
          </div>
        </template>
      </bk-table-column>
    </bk-table>
  </bk-loading>
  <edit-config
    v-model:show="editPanelShow"
    :config-id="activeConfig"
    :bk-biz-id="props.bkBizId"
    :app-id="props.appId"
    @confirm="getListData" />
  <VersionDiff
    v-model:show="isDiffPanelShow"
    :current-version="versionData"
    :current-config="diffConfig" />
</template>
<style lang="scss" scoped>
  .operate-action-btns {
    .bk-button:not(:last-of-type) {
      margin-right: 8px;
    }
  }
</style>